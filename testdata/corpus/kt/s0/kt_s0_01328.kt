package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step10940")
fun step10940(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
object Registry10941 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10941 = "registry-72"
    }
}
object Registry10942 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10942 = "registry-79"
    }
}
fun <T, R> List<T>.mapChecked10943(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked10944(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome10945 {
    data class Success(val payload: String) : Outcome10945()
    data object Failure : Outcome10945()
}

fun describe10945(outcome: Outcome10945): String = when (outcome) {
    is Outcome10945.Success -> outcome.payload
    Outcome10945.Failure -> "failed"
}
