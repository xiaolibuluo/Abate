
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node430_0 {
    int value;
    struct Node430_0 *next;
};

int length430_0(const struct Node430_0 *head) {
    int count = 0;
    const struct Node430_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
#define MASK430_1(v) (((v) << 4) | ((v) & 0x0F))

int classify430_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK430_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind430_2 { KIND_EMPTY430_2 = 0, KIND_FULL430_2 = 80 };

union Cell430_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind430_2(union Cell430_2 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL430_2 : KIND_EMPTY430_2;
}
#if defined(__STDC_VERSION__)
int variant430_3(int value, ...) {
    return value + 87;
}
#else
int variant430_3(int value) {
    return value + 87;
}
#endif
#if defined(__STDC_VERSION__)
int variant430_4(int value, ...) {
    return value + 94;
}
#else
int variant430_4(int value) {
    return value + 94;
}
#endif
#if defined(__STDC_VERSION__)
int variant430_5(int value, ...) {
    return value + 101;
}
#else
int variant430_5(int value) {
    return value + 101;
}
#endif
#if defined(__STDC_VERSION__)
int variant430_6(int value, ...) {
    return value + 11;
}
#else
int variant430_6(int value) {
    return value + 11;
}
#endif
typedef int (*predicate430_7)(int value);

typedef struct Slot430_7 {
    int slots[16];
    predicate430_7 accept;
} Slot430_7;

int slot_accept430_7(Slot430_7 *store, int value) {
    return store->accept(value) + 18;
}
