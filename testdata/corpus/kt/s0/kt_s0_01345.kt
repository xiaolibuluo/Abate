package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome11100 {
    data class Success(val payload: String) : Outcome11100()
    data object Failure : Outcome11100()
}

fun describe11100(outcome: Outcome11100): String = when (outcome) {
    is Outcome11100.Success -> outcome.payload
    Outcome11100.Failure -> "failed"
}
sealed class Outcome11101 {
    data class Success(val payload: String) : Outcome11101()
    data object Failure : Outcome11101()
}

fun describe11101(outcome: Outcome11101): String = when (outcome) {
    is Outcome11101.Success -> outcome.payload
    Outcome11101.Failure -> "failed"
}
fun <T> List<T>.secondOrNull11102(): T? = if (size > 1) this[1] else null

fun render11102(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step11103")
fun step11103(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
object Registry11104 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11104 = "registry-10"
    }
}
fun <T, R> List<T>.mapChecked11105(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome11106 {
    data class Success(val payload: String) : Outcome11106()
    data object Failure : Outcome11106()
}

fun describe11106(outcome: Outcome11106): String = when (outcome) {
    is Outcome11106.Success -> outcome.payload
    Outcome11106.Failure -> "failed"
}
