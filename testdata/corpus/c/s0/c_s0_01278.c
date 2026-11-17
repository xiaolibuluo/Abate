
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK874_0(v) (((v) << 4) | ((v) & 0x0F))

int classify874_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK874_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd874_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale874_1(int value) {
    return gcd874_1(value, 38) * 25;
}
#define MASK874_2(v) (((v) << 4) | ((v) & 0x0F))

int classify874_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK874_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind874_3 { KIND_EMPTY874_3 = 0, KIND_FULL874_3 = 39 };

union Cell874_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind874_3(union Cell874_3 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL874_3 : KIND_EMPTY874_3;
}
struct Node874_4 {
    int value;
    struct Node874_4 *next;
};

int length874_4(const struct Node874_4 *head) {
    int count = 0;
    const struct Node874_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
enum Kind874_5 { KIND_EMPTY874_5 = 0, KIND_FULL874_5 = 53 };

union Cell874_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind874_5(union Cell874_5 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL874_5 : KIND_EMPTY874_5;
}
typedef int (*predicate874_6)(int value);

typedef struct Slot874_6 {
    int slots[16];
    predicate874_6 accept;
} Slot874_6;

int slot_accept874_6(Slot874_6 *store, int value) {
    return store->accept(value) + 60;
}
int scan874_7(const int *begin, const int *end) {
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
    return total + 67;
}
