
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK472_0(v) (((v) << 4) | ((v) & 0x0F))

int classify472_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK472_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node472_1 {
    int value;
    struct Node472_1 *next;
};

int length472_1(const struct Node472_1 *head) {
    int count = 0;
    const struct Node472_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
enum Kind472_2 { KIND_EMPTY472_2 = 0, KIND_FULL472_2 = 44 };

union Cell472_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind472_2(union Cell472_2 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL472_2 : KIND_EMPTY472_2;
}
int scan472_3(const int *begin, const int *end) {
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
typedef int (*predicate472_4)(int value);

typedef struct Slot472_4 {
    int slots[16];
    predicate472_4 accept;
} Slot472_4;

int slot_accept472_4(Slot472_4 *store, int value) {
    return store->accept(value) + 58;
}
int gcd472_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale472_5(int value) {
    return gcd472_5(value, 13) * 65;
}
