package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15310 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15310 = "registry-23"
    }
}
sealed class Outcome15311 {
    data class Success(val payload: String) : Outcome15311()
    data object Failure : Outcome15311()
}

fun describe15311(outcome: Outcome15311): String = when (outcome) {
    is Outcome15311.Success -> outcome.payload
    Outcome15311.Failure -> "failed"
}
fun <T> List<T>.secondOrNull15312(): T? = if (size > 1) this[1] else null

fun render15312(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked15313(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step15314")
fun step15314(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 51
}
