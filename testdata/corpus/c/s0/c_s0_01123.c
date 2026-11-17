
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK769_0(v) (((v) << 4) | ((v) & 0x0F))

int classify769_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK769_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node769_1 {
    int value;
    struct Node769_1 *next;
};

int length769_1(const struct Node769_1 *head) {
    int count = 0;
    const struct Node769_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
#define MASK769_2(v) (((v) << 4) | ((v) & 0x0F))

int classify769_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK769_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind769_3 { KIND_EMPTY769_3 = 0, KIND_FULL769_3 = 32 };

union Cell769_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind769_3(union Cell769_3 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL769_3 : KIND_EMPTY769_3;
}
typedef int (*predicate769_4)(int value);

typedef struct Slot769_4 {
    int slots[16];
    predicate769_4 accept;
} Slot769_4;

int slot_accept769_4(Slot769_4 *store, int value) {
    return store->accept(value) + 39;
}
#if defined(__STDC_VERSION__)
int variant769_5(int value, ...) {
    return value + 46;
}
#else
int variant769_5(int value) {
    return value + 46;
}
#endif
struct Node769_6 {
    int value;
    struct Node769_6 *next;
};

int length769_6(const struct Node769_6 *head) {
    int count = 0;
    const struct Node769_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
