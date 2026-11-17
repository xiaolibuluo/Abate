package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21760 {
    data class Success(val payload: String) : Outcome21760()
    data object Failure : Outcome21760()
}

fun describe21760(outcome: Outcome21760): String = when (outcome) {
    is Outcome21760.Success -> outcome.payload
    Outcome21760.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21761(): T? = if (size > 1) this[1] else null

fun render21761(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull21762(): T? = if (size > 1) this[1] else null

fun render21762(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome21763 {
    data class Success(val payload: String) : Outcome21763()
    data object Failure : Outcome21763()
}

fun describe21763(outcome: Outcome21763): String = when (outcome) {
    is Outcome21763.Success -> outcome.payload
    Outcome21763.Failure -> "failed"
}
data class Pair21764<A, B>(val first: A, val second: B)

fun combine21764(left: Int, right: Int): Pair21764<Int, Int> {
    return Pair21764(left + 94, right + 35)
}
data class Pair21765<A, B>(val first: A, val second: B)

fun combine21765(left: Int, right: Int): Pair21765<Int, Int> {
    return Pair21765(left + 101, right + 46)
}
@JvmName("step21766")
fun step21766(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
