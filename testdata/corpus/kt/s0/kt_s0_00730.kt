package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6010")
fun step6010(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
data class Pair6011<A, B>(val first: A, val second: B)

fun combine6011(left: Int, right: Int): Pair6011<Int, Int> {
    return Pair6011(left + 65, right + 18)
}
fun <T, R> List<T>.mapChecked6012(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome6013 {
    data class Success(val payload: String) : Outcome6013()
    data object Failure : Outcome6013()
}

fun describe6013(outcome: Outcome6013): String = when (outcome) {
    is Outcome6013.Success -> outcome.payload
    Outcome6013.Failure -> "failed"
}
