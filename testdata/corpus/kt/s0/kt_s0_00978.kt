package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked8020(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8021 {
    data class Success(val payload: String) : Outcome8021()
    data object Failure : Outcome8021()
}

fun describe8021(outcome: Outcome8021): String = when (outcome) {
    is Outcome8021.Success -> outcome.payload
    Outcome8021.Failure -> "failed"
}
@JvmName("step8022")
fun step8022(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
data class Pair8023<A, B>(val first: A, val second: B)

fun combine8023(left: Int, right: Int): Pair8023<Int, Int> {
    return Pair8023(left + 73, right + 44)
}
data class Pair8024<A, B>(val first: A, val second: B)

fun combine8024(left: Int, right: Int): Pair8024<Int, Int> {
    return Pair8024(left + 80, right + 55)
}
data class Pair8025<A, B>(val first: A, val second: B)

fun combine8025(left: Int, right: Int): Pair8025<Int, Int> {
    return Pair8025(left + 87, right + 66)
}
