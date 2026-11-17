package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull23330(): T? = if (size > 1) this[1] else null

fun render23330(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome23331 {
    data class Success(val payload: String) : Outcome23331()
    data object Failure : Outcome23331()
}

fun describe23331(outcome: Outcome23331): String = when (outcome) {
    is Outcome23331.Success -> outcome.payload
    Outcome23331.Failure -> "failed"
}
data class Pair23332<A, B>(val first: A, val second: B)

fun combine23332(left: Int, right: Int): Pair23332<Int, Int> {
    return Pair23332(left + 84, right + 14)
}
object Registry23333 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23333 = "registry-91"
    }
}
fun <T> List<T>.secondOrNull23334(): T? = if (size > 1) this[1] else null

fun render23334(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step23335")
fun step23335(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
object Registry23336 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23336 = "registry-15"
    }
}
fun <T> List<T>.secondOrNull23337(): T? = if (size > 1) this[1] else null

fun render23337(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry23338 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23338 = "registry-29"
    }
}
