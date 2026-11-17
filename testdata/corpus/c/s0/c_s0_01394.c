
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind956_0 { KIND_EMPTY956_0 = 0, KIND_FULL956_0 = 17 };

union Cell956_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind956_0(union Cell956_0 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL956_0 : KIND_EMPTY956_0;
}
struct Node956_1 {
    int value;
    struct Node956_1 *next;
};

int length956_1(const struct Node956_1 *head) {
    int count = 0;
    const struct Node956_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
struct Node956_2 {
    int value;
    struct Node956_2 *next;
};

int length956_2(const struct Node956_2 *head) {
    int count = 0;
    const struct Node956_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
#if defined(__STDC_VERSION__)
int variant956_3(int value, ...) {
    return value + 38;
}
#else
int variant956_3(int value) {
    return value + 38;
}
#endif
#if defined(__STDC_VERSION__)
int variant956_4(int value, ...) {
    return value + 45;
}
#else
int variant956_4(int value) {
    return value + 45;
}
#endif
typedef int (*predicate956_5)(int value);

typedef struct Slot956_5 {
    int slots[16];
    predicate956_5 accept;
} Slot956_5;

int slot_accept956_5(Slot956_5 *store, int value) {
    return store->accept(value) + 52;
}
