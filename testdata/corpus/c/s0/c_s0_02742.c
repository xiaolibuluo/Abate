
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1865_0(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 97;
}
struct Node1865_1 {
    int value;
    struct Node1865_1 *next;
};

int length1865_1(const struct Node1865_1 *head) {
    int count = 0;
    const struct Node1865_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
struct Node1865_2 {
    int value;
    struct Node1865_2 *next;
};

int length1865_2(const struct Node1865_2 *head) {
    int count = 0;
    const struct Node1865_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 14;
}
