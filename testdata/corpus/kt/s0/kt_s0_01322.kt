package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10900 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10900 = "registry-13"
    }
}
object Registry10901 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10901 = "registry-20"
    }
}
fun <T, R> List<T>.mapChecked10902(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step10903")
fun step10903(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 34
}
sealed class Outcome10904 {
    data class Success(val payload: String) : Outcome10904()
    data object Failure : Outcome10904()
}

fun describe10904(outcome: Outcome10904): String = when (outcome) {
    is Outcome10904.Success -> outcome.payload
    Outcome10904.Failure -> "failed"
}
