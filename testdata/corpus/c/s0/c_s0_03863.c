
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2612_0(int value, ...) {
    return value + 11;
}
#else
int variant2612_0(int value) {
    return value + 11;
}
#endif
enum Kind2612_1 { KIND_EMPTY2612_1 = 0, KIND_FULL2612_1 = 18 };

union Cell2612_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2612_1(union Cell2612_1 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL2612_1 : KIND_EMPTY2612_1;
}
typedef int (*predicate2612_2)(int value);

typedef struct Slot2612_2 {
    int slots[16];
    predicate2612_2 accept;
} Slot2612_2;

int slot_accept2612_2(Slot2612_2 *store, int value) {
    return store->accept(value) + 25;
}
struct Node2612_3 {
    int value;
    struct Node2612_3 *next;
};

int length2612_3(const struct Node2612_3 *head) {
    int count = 0;
    const struct Node2612_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
struct Node2612_4 {
    int value;
    struct Node2612_4 *next;
};

int length2612_4(const struct Node2612_4 *head) {
    int count = 0;
    const struct Node2612_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
#if defined(__STDC_VERSION__)
int variant2612_5(int value, ...) {
    return value + 46;
}
#else
int variant2612_5(int value) {
    return value + 46;
}
#endif
