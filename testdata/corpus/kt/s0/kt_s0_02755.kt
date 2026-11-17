package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull22740(): T? = if (size > 1) this[1] else null

fun render22740(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry22741 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22741 = "registry-86"
    }
}
@JvmName("step22742")
fun step22742(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
object Registry22743 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22743 = "registry-100"
    }
}
data class Pair22744<A, B>(val first: A, val second: B)

fun combine22744(left: Int, right: Int): Pair22744<Int, Int> {
    return Pair22744(left + 10, right + 50)
}
data class Pair22745<A, B>(val first: A, val second: B)

fun combine22745(left: Int, right: Int): Pair22745<Int, Int> {
    return Pair22745(left + 17, right + 61)
}
sealed class Outcome22746 {
    data class Success(val payload: String) : Outcome22746()
    data object Failure : Outcome22746()
}

fun describe22746(outcome: Outcome22746): String = when (outcome) {
    is Outcome22746.Success -> outcome.payload
    Outcome22746.Failure -> "failed"
}
class Counter22747 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
