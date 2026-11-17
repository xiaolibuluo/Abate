package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step4720")
fun step4720(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
sealed class Outcome4721 {
    data class Success(val payload: String) : Outcome4721()
    data object Failure : Outcome4721()
}

fun describe4721(outcome: Outcome4721): String = when (outcome) {
    is Outcome4721.Success -> outcome.payload
    Outcome4721.Failure -> "failed"
}
data class Pair4722<A, B>(val first: A, val second: B)

fun combine4722(left: Int, right: Int): Pair4722<Int, Int> {
    return Pair4722(left + 44, right + 41)
}
sealed class Outcome4723 {
    data class Success(val payload: String) : Outcome4723()
    data object Failure : Outcome4723()
}

fun describe4723(outcome: Outcome4723): String = when (outcome) {
    is Outcome4723.Success -> outcome.payload
    Outcome4723.Failure -> "failed"
}
object Registry4724 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4724 = "registry-58"
    }
}
fun <T, R> List<T>.mapChecked4725(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome4726 {
    data class Success(val payload: String) : Outcome4726()
    data object Failure : Outcome4726()
}

fun describe4726(outcome: Outcome4726): String = when (outcome) {
    is Outcome4726.Success -> outcome.payload
    Outcome4726.Failure -> "failed"
}
