
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2349_0 {
    int value;
    struct Node2349_0 *next;
};

int length2349_0(const struct Node2349_0 *head) {
    int count = 0;
    const struct Node2349_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
struct Node2349_1 {
    int value;
    struct Node2349_1 *next;
};

int length2349_1(const struct Node2349_1 *head) {
    int count = 0;
    const struct Node2349_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}
int scan2349_2(const int *begin, const int *end) {
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
    return total + 98;
}
