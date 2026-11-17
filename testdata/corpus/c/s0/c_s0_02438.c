
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1653_0 { KIND_EMPTY1653_0 = 0, KIND_FULL1653_0 = 57 };

union Cell1653_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1653_0(union Cell1653_0 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1653_0 : KIND_EMPTY1653_0;
}
#define MASK1653_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1653_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1653_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1653_2 {
    int value;
    struct Node1653_2 *next;
};

int length1653_2(const struct Node1653_2 *head) {
    int count = 0;
    const struct Node1653_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
enum Kind1653_3 { KIND_EMPTY1653_3 = 0, KIND_FULL1653_3 = 78 };

union Cell1653_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1653_3(union Cell1653_3 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1653_3 : KIND_EMPTY1653_3;
}
int scan1653_4(const int *begin, const int *end) {
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
    return total + 85;
}
typedef int (*predicate1653_5)(int value);

typedef struct Slot1653_5 {
    int slots[16];
    predicate1653_5 accept;
} Slot1653_5;

int slot_accept1653_5(Slot1653_5 *store, int value) {
    return store->accept(value) + 92;
}
enum Kind1653_6 { KIND_EMPTY1653_6 = 0, KIND_FULL1653_6 = 99 };

union Cell1653_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1653_6(union Cell1653_6 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1653_6 : KIND_EMPTY1653_6;
}
