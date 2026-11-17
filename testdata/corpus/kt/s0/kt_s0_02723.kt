package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22480<A, B>(val first: A, val second: B)

fun combine22480(left: Int, right: Int): Pair22480<Int, Int> {
    return Pair22480(left + 32, right + 68)
}
@JvmName("step22481")
fun step22481(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
data class Pair22482<A, B>(val first: A, val second: B)

fun combine22482(left: Int, right: Int): Pair22482<Int, Int> {
    return Pair22482(left + 46, right + 29)
}
