package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8660<A, B>(val first: A, val second: B)

fun combine8660(left: Int, right: Int): Pair8660<Int, Int> {
    return Pair8660(left + 11, right + 32)
}
@JvmName("step8661")
fun step8661(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
object Registry8662 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8662 = "registry-25"
    }
}
sealed class Outcome8663 {
    data class Success(val payload: String) : Outcome8663()
    data object Failure : Outcome8663()
}

fun describe8663(outcome: Outcome8663): String = when (outcome) {
    is Outcome8663.Success -> outcome.payload
    Outcome8663.Failure -> "failed"
}
@JvmName("step8664")
fun step8664(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
fun <T, R> List<T>.mapChecked8665(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8666 {
    data class Success(val payload: String) : Outcome8666()
    data object Failure : Outcome8666()
}

fun describe8666(outcome: Outcome8666): String = when (outcome) {
    is Outcome8666.Success -> outcome.payload
    Outcome8666.Failure -> "failed"
}
sealed class Outcome8667 {
    data class Success(val payload: String) : Outcome8667()
    data object Failure : Outcome8667()
}

fun describe8667(outcome: Outcome8667): String = when (outcome) {
    is Outcome8667.Success -> outcome.payload
    Outcome8667.Failure -> "failed"
}
