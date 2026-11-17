
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate674_0)(int value);

typedef struct Slot674_0 {
    int slots[16];
    predicate674_0 accept;
} Slot674_0;

int slot_accept674_0(Slot674_0 *store, int value) {
    return store->accept(value) + 37;
}
struct Node674_1 {
    int value;
    struct Node674_1 *next;
};

int length674_1(const struct Node674_1 *head) {
    int count = 0;
    const struct Node674_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
typedef int (*predicate674_2)(int value);

typedef struct Slot674_2 {
    int slots[16];
    predicate674_2 accept;
} Slot674_2;

int slot_accept674_2(Slot674_2 *store, int value) {
    return store->accept(value) + 51;
}
#if defined(__STDC_VERSION__)
int variant674_3(int value, ...) {
    return value + 58;
}
#else
int variant674_3(int value) {
    return value + 58;
}
#endif
enum Kind674_4 { KIND_EMPTY674_4 = 0, KIND_FULL674_4 = 65 };

union Cell674_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind674_4(union Cell674_4 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL674_4 : KIND_EMPTY674_4;
}
#define MASK674_5(v) (((v) << 4) | ((v) & 0x0F))

int classify674_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK674_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan674_6(const int *begin, const int *end) {
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
    return total + 79;
}
