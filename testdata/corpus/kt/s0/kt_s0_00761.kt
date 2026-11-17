package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6280")
fun step6280(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
sealed class Outcome6281 {
    data class Success(val payload: String) : Outcome6281()
    data object Failure : Outcome6281()
}

fun describe6281(outcome: Outcome6281): String = when (outcome) {
    is Outcome6281.Success -> outcome.payload
    Outcome6281.Failure -> "failed"
}
data class Pair6282<A, B>(val first: A, val second: B)

fun combine6282(left: Int, right: Int): Pair6282<Int, Int> {
    return Pair6282(left + 35, right + 35)
}
@JvmName("step6283")
fun step6283(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
fun <T> List<T>.secondOrNull6284(): T? = if (size > 1) this[1] else null

fun render6284(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6285 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6285 = "registry-56"
    }
}
@JvmName("step6286")
fun step6286(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
object Registry6287 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6287 = "registry-70"
    }
}
class Counter6288 {
    var current: Int by Delegates.observable(77) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
