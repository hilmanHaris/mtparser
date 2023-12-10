package mtparser

import "errors"

func (s *Parser) scanBody() error {
	var fld Field
	var p rune
	mp := map[string]Node{}
	bin := len(s.Blocks)

	s.blk.Val = []Field{}

	if s.Scan() != '\n' {
		return errors.New(s.ErrMessage('\n', true))
	}

	for i2 := 0; s.Peek() != '}'; i2++ {

		if s.Scan() != ':' {
			return errors.New(s.ErrMessage(':', true))
		}

		s.Scan()
		fld.Key = s.TokenText()

		if s.Scan() != ':' {
			return errors.New(s.ErrMessage(':', true))
		}

		fld.Val = ""

		for {
			t := s.Scan()
			p = s.Peek()

			if t == -1 {
				break
			}
			if p == '-' && t == '\n' {
				break
			}
			if t == '\n' && p == ':' {
				break
			}

			fld.Val += s.TokenText()
		}

		mp[fld.Key] = Node{
			Val: fld.Val,
			Blk: bin,
			Ind: i2,
		}
		s.blk.Val = append(s.blk.Val.([]Field), *&fld)

		if p == '-' {
			s.Scan()
			break
		}
	}

	s.Map[s.blk.Key] = mp
	return nil
}
