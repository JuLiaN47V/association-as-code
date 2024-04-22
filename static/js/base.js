function changeText(self, show_more, show_less){
    if (show_more == null){
        show_more = "Show more"    
    }
    if (show_less == null){
        show_less = "Show less"    
    }
    if (self.innerText == show_more){
        self.innerText = show_less
    } else {
        self.innerText = show_more
    }
}