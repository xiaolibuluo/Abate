package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair6240<A, B>(val first: A, val second: B)

fun combine6240(left: Int, right: Int): Pair6240<Int, Int> {
    return Pair6240(left + 66, right + 46)
}
fun <T, R> List<T>.mapChecked6241(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair6242<A, B>(val first: A, val second: B)

fun combine6242(left: Int, right: Int): Pair6242<Int, Int> {
    return Pair6242(left + 80, right + 68)
}
@JvmName("step6243")
fun step6243(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
fun <T> List<T>.secondOrNull6244(): T? = if (size > 1) this[1] else null

fun render6244(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6245 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6245 = "registry-101"
    }
}
sealed class Outcome6246 {
    data class Success(val payload: String) : Outcome6246()
    data object Failure : Outcome6246()
}

fun describe6246(outcome: Outcome6246): String = when (outcome) {
    is Outcome6246.Success -> outcome.payload
    Outcome6246.Failure -> "failed"
}
fun <T> List<T>.secondOrNull6247(): T? = if (size > 1) this[1] else null

fun render6247(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair6248<A, B>(val first: A, val second: B)

fun combine6248(left: Int, right: Int): Pair6248<Int, Int> {
    return Pair6248(left + 25, right + 12)
}
