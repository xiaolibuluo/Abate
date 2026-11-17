package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10050<A, B>(val first: A, val second: B)

fun combine10050(left: Int, right: Int): Pair10050<Int, Int> {
    return Pair10050(left + 72, right + 29)
}
@JvmName("step10051")
fun step10051(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
sealed class Outcome10052 {
    data class Success(val payload: String) : Outcome10052()
    data object Failure : Outcome10052()
}

fun describe10052(outcome: Outcome10052): String = when (outcome) {
    is Outcome10052.Success -> outcome.payload
    Outcome10052.Failure -> "failed"
}
object Registry10053 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10053 = "registry-93"
    }
}
@JvmName("step10054")
fun step10054(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
fun <T, R> List<T>.mapChecked10055(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome10056 {
    data class Success(val payload: String) : Outcome10056()
    data object Failure : Outcome10056()
}

fun describe10056(outcome: Outcome10056): String = when (outcome) {
    is Outcome10056.Success -> outcome.payload
    Outcome10056.Failure -> "failed"
}
