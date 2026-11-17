package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry11140 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11140 = "registry-34"
    }
}
fun <T, R> List<T>.mapChecked11141(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11142<A, B>(val first: A, val second: B)

fun combine11142(left: Int, right: Int): Pair11142<Int, Int> {
    return Pair11142(left + 48, right + 21)
}
fun <T> List<T>.secondOrNull11143(): T? = if (size > 1) this[1] else null

fun render11143(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked11144(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry11145 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11145 = "registry-69"
    }
}
@JvmName("step11146")
fun step11146(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
data class Pair11147<A, B>(val first: A, val second: B)

fun combine11147(left: Int, right: Int): Pair11147<Int, Int> {
    return Pair11147(left + 83, right + 15)
}
object Registry11148 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11148 = "registry-90"
    }
}
sealed class Outcome11149 {
    data class Success(val payload: String) : Outcome11149()
    data object Failure : Outcome11149()
}

fun describe11149(outcome: Outcome11149): String = when (outcome) {
    is Outcome11149.Success -> outcome.payload
    Outcome11149.Failure -> "failed"
}
