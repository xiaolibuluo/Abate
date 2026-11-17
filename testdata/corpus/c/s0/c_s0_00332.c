
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node228_0 {
    int value;
    struct Node228_0 *next;
};

int length228_0(const struct Node228_0 *head) {
    int count = 0;
    const struct Node228_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
int scan228_1(const int *begin, const int *end) {
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
    return total + 66;
}
enum Kind228_2 { KIND_EMPTY228_2 = 0, KIND_FULL228_2 = 73 };

union Cell228_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind228_2(union Cell228_2 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL228_2 : KIND_EMPTY228_2;
}
struct Node228_3 {
    int value;
    struct Node228_3 *next;
};

int length228_3(const struct Node228_3 *head) {
    int count = 0;
    const struct Node228_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}
typedef int (*predicate228_4)(int value);

typedef struct Slot228_4 {
    int slots[16];
    predicate228_4 accept;
} Slot228_4;

int slot_accept228_4(Slot228_4 *store, int value) {
    return store->accept(value) + 87;
}
