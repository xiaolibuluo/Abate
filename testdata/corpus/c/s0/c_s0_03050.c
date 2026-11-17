
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2059_0 {
    int value;
    struct Node2059_0 *next;
};

int length2059_0(const struct Node2059_0 *head) {
    int count = 0;
    const struct Node2059_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
struct Node2059_1 {
    int value;
    struct Node2059_1 *next;
};

int length2059_1(const struct Node2059_1 *head) {
    int count = 0;
    const struct Node2059_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
enum Kind2059_2 { KIND_EMPTY2059_2 = 0, KIND_FULL2059_2 = 14 };

union Cell2059_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2059_2(union Cell2059_2 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL2059_2 : KIND_EMPTY2059_2;
}
