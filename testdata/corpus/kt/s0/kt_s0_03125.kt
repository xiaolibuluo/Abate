package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry25640 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25640 = "registry-66"
    }
}
sealed class Outcome25641 {
    data class Success(val payload: String) : Outcome25641()
    data object Failure : Outcome25641()
}

fun describe25641(outcome: Outcome25641): String = when (outcome) {
    is Outcome25641.Success -> outcome.payload
    Outcome25641.Failure -> "failed"
}
data class Pair25642<A, B>(val first: A, val second: B)

fun combine25642(left: Int, right: Int): Pair25642<Int, Int> {
    return Pair25642(left + 80, right + 45)
}
object Registry25643 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25643 = "registry-87"
    }
}
object Registry25644 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25644 = "registry-94"
    }
}
fun <T> List<T>.secondOrNull25645(): T? = if (size > 1) this[1] else null

fun render25645(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step25646")
fun step25646(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
