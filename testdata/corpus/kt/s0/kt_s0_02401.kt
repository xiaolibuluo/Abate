package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19720")
fun step19720(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
fun <T, R> List<T>.mapChecked19721(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step19722")
fun step19722(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
sealed class Outcome19723 {
    data class Success(val payload: String) : Outcome19723()
    data object Failure : Outcome19723()
}

fun describe19723(outcome: Outcome19723): String = when (outcome) {
    is Outcome19723.Success -> outcome.payload
    Outcome19723.Failure -> "failed"
}
object Registry19724 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19724 = "registry-61"
    }
}
fun <T> List<T>.secondOrNull19725(): T? = if (size > 1) this[1] else null

fun render19725(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
