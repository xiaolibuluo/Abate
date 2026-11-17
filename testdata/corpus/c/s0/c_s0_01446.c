
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node984_0 {
    int value;
    struct Node984_0 *next;
};

int length984_0(const struct Node984_0 *head) {
    int count = 0;
    const struct Node984_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 90;
}
enum Kind984_1 { KIND_EMPTY984_1 = 0, KIND_FULL984_1 = 97 };

union Cell984_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind984_1(union Cell984_1 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL984_1 : KIND_EMPTY984_1;
}
