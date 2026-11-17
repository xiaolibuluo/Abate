
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node382_0 {
    int value;
    struct Node382_0 *next;
};

int length382_0(const struct Node382_0 *head) {
    int count = 0;
    const struct Node382_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
int gcd382_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale382_1(int value) {
    return gcd382_1(value, 10) * 31;
}
typedef int (*predicate382_2)(int value);

typedef struct Slot382_2 {
    int slots[16];
    predicate382_2 accept;
} Slot382_2;

int slot_accept382_2(Slot382_2 *store, int value) {
    return store->accept(value) + 38;
}
#define MASK382_3(v) (((v) << 4) | ((v) & 0x0F))

int classify382_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK382_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node382_4 {
    int value;
    struct Node382_4 *next;
};

int length382_4(const struct Node382_4 *head) {
    int count = 0;
    const struct Node382_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
int scan382_5(const int *begin, const int *end) {
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
    return total + 59;
}
enum Kind382_6 { KIND_EMPTY382_6 = 0, KIND_FULL382_6 = 66 };

union Cell382_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind382_6(union Cell382_6 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL382_6 : KIND_EMPTY382_6;
}
