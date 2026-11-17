package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked26270(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked26271(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step26272")
fun step26272(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
sealed class Outcome26273 {
    data class Success(val payload: String) : Outcome26273()
    data object Failure : Outcome26273()
}

fun describe26273(outcome: Outcome26273): String = when (outcome) {
    is Outcome26273.Success -> outcome.payload
    Outcome26273.Failure -> "failed"
}
fun <T> List<T>.secondOrNull26274(): T? = if (size > 1) this[1] else null

fun render26274(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair26275<A, B>(val first: A, val second: B)

fun combine26275(left: Int, right: Int): Pair26275<Int, Int> {
    return Pair26275(left + 47, right + 31)
}
fun <T> List<T>.secondOrNull26276(): T? = if (size > 1) this[1] else null

fun render26276(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull26277(): T? = if (size > 1) this[1] else null

fun render26277(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull26278(): T? = if (size > 1) this[1] else null

fun render26278(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry26279 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26279 = "registry-75"
    }
}
