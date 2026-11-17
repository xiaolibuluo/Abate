
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2442_0 {
    int value;
    struct Node2442_0 *next;
};

int length2442_0(const struct Node2442_0 *head) {
    int count = 0;
    const struct Node2442_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
#if defined(__STDC_VERSION__)
int variant2442_1(int value, ...) {
    return value + 39;
}
#else
int variant2442_1(int value) {
    return value + 39;
}
#endif
struct Node2442_2 {
    int value;
    struct Node2442_2 *next;
};

int length2442_2(const struct Node2442_2 *head) {
    int count = 0;
    const struct Node2442_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
