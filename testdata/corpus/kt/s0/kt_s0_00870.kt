package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull7130(): T? = if (size > 1) this[1] else null

fun render7130(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry7131 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7131 = "registry-66"
    }
}
@JvmName("step7132")
fun step7132(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 73
}
sealed class Outcome7133 {
    data class Success(val payload: String) : Outcome7133()
    data object Failure : Outcome7133()
}

fun describe7133(outcome: Outcome7133): String = when (outcome) {
    is Outcome7133.Success -> outcome.payload
    Outcome7133.Failure -> "failed"
}
sealed class Outcome7134 {
    data class Success(val payload: String) : Outcome7134()
    data object Failure : Outcome7134()
}

fun describe7134(outcome: Outcome7134): String = when (outcome) {
    is Outcome7134.Success -> outcome.payload
    Outcome7134.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked7135(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
