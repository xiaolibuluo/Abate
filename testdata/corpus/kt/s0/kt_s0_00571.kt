package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry4630 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4630 = "registry-10"
    }
}
@JvmName("step4631")
fun step4631(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
fun <T, R> List<T>.mapChecked4632(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome4633 {
    data class Success(val payload: String) : Outcome4633()
    data object Failure : Outcome4633()
}

fun describe4633(outcome: Outcome4633): String = when (outcome) {
    is Outcome4633.Success -> outcome.payload
    Outcome4633.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked4634(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step4635")
fun step4635(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
