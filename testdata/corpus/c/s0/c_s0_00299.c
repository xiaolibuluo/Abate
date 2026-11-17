
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node207_0 {
    int value;
    struct Node207_0 *next;
};

int length207_0(const struct Node207_0 *head) {
    int count = 0;
    const struct Node207_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 77;
}
#if defined(__STDC_VERSION__)
int variant207_1(int value, ...) {
    return value + 84;
}
#else
int variant207_1(int value) {
    return value + 84;
}
#endif
#if defined(__STDC_VERSION__)
int variant207_2(int value, ...) {
    return value + 91;
}
#else
int variant207_2(int value) {
    return value + 91;
}
#endif
