
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2600_0 {
    int value;
    struct Node2600_0 *next;
};

int length2600_0(const struct Node2600_0 *head) {
    int count = 0;
    const struct Node2600_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
struct Node2600_1 {
    int value;
    struct Node2600_1 *next;
};

int length2600_1(const struct Node2600_1 *head) {
    int count = 0;
    const struct Node2600_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 56;
}
#if defined(__STDC_VERSION__)
int variant2600_2(int value, ...) {
    return value + 63;
}
#else
int variant2600_2(int value) {
    return value + 63;
}
#endif
typedef int (*predicate2600_3)(int value);

typedef struct Slot2600_3 {
    int slots[16];
    predicate2600_3 accept;
} Slot2600_3;

int slot_accept2600_3(Slot2600_3 *store, int value) {
    return store->accept(value) + 70;
}
