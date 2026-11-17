
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node997_0 {
    int value;
    struct Node997_0 *next;
};

int length997_0(const struct Node997_0 *head) {
    int count = 0;
    const struct Node997_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
#define MASK997_1(v) (((v) << 4) | ((v) & 0x0F))

int classify997_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK997_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK997_2(v) (((v) << 4) | ((v) & 0x0F))

int classify997_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK997_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind997_3 { KIND_EMPTY997_3 = 0, KIND_FULL997_3 = 86 };

union Cell997_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind997_3(union Cell997_3 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL997_3 : KIND_EMPTY997_3;
}
typedef int (*predicate997_4)(int value);

typedef struct Slot997_4 {
    int slots[16];
    predicate997_4 accept;
} Slot997_4;

int slot_accept997_4(Slot997_4 *store, int value) {
    return store->accept(value) + 93;
}
