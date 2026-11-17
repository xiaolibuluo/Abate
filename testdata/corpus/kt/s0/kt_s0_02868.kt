package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23670<A, B>(val first: A, val second: B)

fun combine23670(left: Int, right: Int): Pair23670<Int, Int> {
    return Pair23670(left + 27, right + 47)
}
data class Pair23671<A, B>(val first: A, val second: B)

fun combine23671(left: Int, right: Int): Pair23671<Int, Int> {
    return Pair23671(left + 34, right + 58)
}
@JvmName("step23672")
fun step23672(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
data class Pair23673<A, B>(val first: A, val second: B)

fun combine23673(left: Int, right: Int): Pair23673<Int, Int> {
    return Pair23673(left + 48, right + 19)
}
data class Pair23674<A, B>(val first: A, val second: B)

fun combine23674(left: Int, right: Int): Pair23674<Int, Int> {
    return Pair23674(left + 55, right + 30)
}
sealed class Outcome23675 {
    data class Success(val payload: String) : Outcome23675()
    data object Failure : Outcome23675()
}

fun describe23675(outcome: Outcome23675): String = when (outcome) {
    is Outcome23675.Success -> outcome.payload
    Outcome23675.Failure -> "failed"
}
data class Pair23676<A, B>(val first: A, val second: B)

fun combine23676(left: Int, right: Int): Pair23676<Int, Int> {
    return Pair23676(left + 69, right + 52)
}
data class Pair23677<A, B>(val first: A, val second: B)

fun combine23677(left: Int, right: Int): Pair23677<Int, Int> {
    return Pair23677(left + 76, right + 63)
}
