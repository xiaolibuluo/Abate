
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node72_0 {
    int value;
    struct Node72_0 *next;
};

int length72_0(const struct Node72_0 *head) {
    int count = 0;
    const struct Node72_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
#if defined(__STDC_VERSION__)
int variant72_1(int value, ...) {
    return value + 75;
}
#else
int variant72_1(int value) {
    return value + 75;
}
#endif
typedef int (*predicate72_2)(int value);

typedef struct Slot72_2 {
    int slots[16];
    predicate72_2 accept;
} Slot72_2;

int slot_accept72_2(Slot72_2 *store, int value) {
    return store->accept(value) + 82;
}
enum Kind72_3 { KIND_EMPTY72_3 = 0, KIND_FULL72_3 = 89 };

union Cell72_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind72_3(union Cell72_3 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL72_3 : KIND_EMPTY72_3;
}
#if defined(__STDC_VERSION__)
int variant72_4(int value, ...) {
    return value + 96;
}
#else
int variant72_4(int value) {
    return value + 96;
}
#endif
