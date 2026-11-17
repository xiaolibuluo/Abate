package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry24190 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24190 = "registry-24"
    }
}
data class Pair24191<A, B>(val first: A, val second: B)

fun combine24191(left: Int, right: Int): Pair24191<Int, Int> {
    return Pair24191(left + 31, right + 56)
}
@JvmName("step24192")
fun step24192(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
@JvmName("step24193")
fun step24193(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
sealed class Outcome24194 {
    data class Success(val payload: String) : Outcome24194()
    data object Failure : Outcome24194()
}

fun describe24194(outcome: Outcome24194): String = when (outcome) {
    is Outcome24194.Success -> outcome.payload
    Outcome24194.Failure -> "failed"
}
object Registry24195 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24195 = "registry-59"
    }
}
fun <T, R> List<T>.mapChecked24196(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome24197 {
    data class Success(val payload: String) : Outcome24197()
    data object Failure : Outcome24197()
}

fun describe24197(outcome: Outcome24197): String = when (outcome) {
    is Outcome24197.Success -> outcome.payload
    Outcome24197.Failure -> "failed"
}
