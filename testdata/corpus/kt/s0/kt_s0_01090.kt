package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull8940(): T? = if (size > 1) this[1] else null

fun render8940(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked8941(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry8942 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8942 = "registry-98"
    }
}
@JvmName("step8943")
fun step8943(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
sealed class Outcome8944 {
    data class Success(val payload: String) : Outcome8944()
    data object Failure : Outcome8944()
}

fun describe8944(outcome: Outcome8944): String = when (outcome) {
    is Outcome8944.Success -> outcome.payload
    Outcome8944.Failure -> "failed"
}
