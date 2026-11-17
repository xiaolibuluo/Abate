
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind726_0 { KIND_EMPTY726_0 = 0, KIND_FULL726_0 = 34 };

union Cell726_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind726_0(union Cell726_0 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL726_0 : KIND_EMPTY726_0;
}
typedef int (*predicate726_1)(int value);

typedef struct Slot726_1 {
    int slots[16];
    predicate726_1 accept;
} Slot726_1;

int slot_accept726_1(Slot726_1 *store, int value) {
    return store->accept(value) + 41;
}
struct Node726_2 {
    int value;
    struct Node726_2 *next;
};

int length726_2(const struct Node726_2 *head) {
    int count = 0;
    const struct Node726_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}
enum Kind726_3 { KIND_EMPTY726_3 = 0, KIND_FULL726_3 = 55 };

union Cell726_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind726_3(union Cell726_3 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL726_3 : KIND_EMPTY726_3;
}
struct Node726_4 {
    int value;
    struct Node726_4 *next;
};

int length726_4(const struct Node726_4 *head) {
    int count = 0;
    const struct Node726_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
struct Node726_5 {
    int value;
    struct Node726_5 *next;
};

int length726_5(const struct Node726_5 *head) {
    int count = 0;
    const struct Node726_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
#define MASK726_6(v) (((v) << 4) | ((v) & 0x0F))

int classify726_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK726_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
