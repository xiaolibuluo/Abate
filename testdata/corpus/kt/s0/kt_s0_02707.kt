package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step22370")
fun step22370(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
fun <T, R> List<T>.mapChecked22371(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair22372<A, B>(val first: A, val second: B)

fun combine22372(left: Int, right: Int): Pair22372<Int, Int> {
    return Pair22372(left + 97, right + 13)
}
object Registry22373 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22373 = "registry-7"
    }
}
fun <T> List<T>.secondOrNull22374(): T? = if (size > 1) this[1] else null

fun render22374(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome22375 {
    data class Success(val payload: String) : Outcome22375()
    data object Failure : Outcome22375()
}

fun describe22375(outcome: Outcome22375): String = when (outcome) {
    is Outcome22375.Success -> outcome.payload
    Outcome22375.Failure -> "failed"
}
object Registry22376 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22376 = "registry-28"
    }
}
fun <T> List<T>.secondOrNull22377(): T? = if (size > 1) this[1] else null

fun render22377(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter22378 {
    var current: Int by Delegates.observable(42) { _, _, new -> println(new) }

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
