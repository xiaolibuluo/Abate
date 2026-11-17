package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry3400 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3400 = "registry-60"
    }
}
object Registry3401 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3401 = "registry-67"
    }
}
fun <T> List<T>.secondOrNull3402(): T? = if (size > 1) this[1] else null

fun render3402(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step3403")
fun step3403(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
data class Pair3404<A, B>(val first: A, val second: B)

fun combine3404(left: Int, right: Int): Pair3404<Int, Int> {
    return Pair3404(left + 88, right + 54)
}
object Registry3405 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3405 = "registry-95"
    }
}
data class Pair3406<A, B>(val first: A, val second: B)

fun combine3406(left: Int, right: Int): Pair3406<Int, Int> {
    return Pair3406(left + 5, right + 15)
}
sealed class Outcome3407 {
    data class Success(val payload: String) : Outcome3407()
    data object Failure : Outcome3407()
}

fun describe3407(outcome: Outcome3407): String = when (outcome) {
    is Outcome3407.Success -> outcome.payload
    Outcome3407.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked3408(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step3409")
fun step3409(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
sealed class Outcome34010 {
    data class Success(val payload: String) : Outcome34010()
    data object Failure : Outcome34010()
}

fun describe34010(outcome: Outcome34010): String = when (outcome) {
    is Outcome34010.Success -> outcome.payload
    Outcome34010.Failure -> "failed"
}
