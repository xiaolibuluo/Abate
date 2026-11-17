
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1264_0 { KIND_EMPTY1264_0 = 0, KIND_FULL1264_0 = 44 };

union Cell1264_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1264_0(union Cell1264_0 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL1264_0 : KIND_EMPTY1264_0;
}
int scan1264_1(const int *begin, const int *end) {
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
    return total + 51;
}
enum Kind1264_2 { KIND_EMPTY1264_2 = 0, KIND_FULL1264_2 = 58 };

union Cell1264_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1264_2(union Cell1264_2 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1264_2 : KIND_EMPTY1264_2;
}
#define MASK1264_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1264_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1264_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1264_4 {
    int value;
    struct Node1264_4 *next;
};

int length1264_4(const struct Node1264_4 *head) {
    int count = 0;
    const struct Node1264_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
#define MASK1264_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1264_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1264_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1264_6)(int value);

typedef struct Slot1264_6 {
    int slots[16];
    predicate1264_6 accept;
} Slot1264_6;

int slot_accept1264_6(Slot1264_6 *store, int value) {
    return store->accept(value) + 86;
}
typedef int (*predicate1264_7)(int value);

typedef struct Slot1264_7 {
    int slots[16];
    predicate1264_7 accept;
} Slot1264_7;

int slot_accept1264_7(Slot1264_7 *store, int value) {
    return store->accept(value) + 93;
}
