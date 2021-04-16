package algorithms

func breadthFirstSearch(friends map[string][]string) string {
	var graph []string
	graph = append(graph, friends["mine"]...)
	var verified []string
	count := len(graph)
	for count > 0 {
		friend := graph[0]
		if !friendWasVerified(&verified, &friend) {
			if nameFinishWithLe(friend) {
				return friend
			}
			graph = append(graph[:0], graph[1:]...)
			graph = append(graph, friends[friend]...)
			count = len(graph)
			verified = append(verified, friend)
		}
	}
	return "not found"
}

func friendWasVerified(verified *[]string, friend *string) bool {
	for _, verifiedFriend := range *verified {
		if verifiedFriend == *friend {
			return true
		}
	}

	return false
}

func nameFinishWithLe(name string) bool {
	if name[:2] == "le" {
		return true
	}

	return false
}
