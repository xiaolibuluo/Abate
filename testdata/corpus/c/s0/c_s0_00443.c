
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant300_0(int value, ...) {
    return value + 25;
}
#else
int variant300_0(int value) {
    return value + 25;
}
#endif
#if defined(__STDC_VERSION__)
int variant300_1(int value, ...) {
    return value + 32;
}
#else
int variant300_1(int value) {
    return value + 32;
}
#endif
struct Node300_2 {
    int value;
    struct Node300_2 *next;
};

int length300_2(const struct Node300_2 *head) {
    int count = 0;
    const struct Node300_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
