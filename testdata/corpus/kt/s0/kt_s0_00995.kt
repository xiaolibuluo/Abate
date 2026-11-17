package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8150 {
    data class Success(val payload: String) : Outcome8150()
    data object Failure : Outcome8150()
}

fun describe8150(outcome: Outcome8150): String = when (outcome) {
    is Outcome8150.Success -> outcome.payload
    Outcome8150.Failure -> "failed"
}
object Registry8151 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8151 = "registry-34"
    }
}
fun <T> List<T>.secondOrNull8152(): T? = if (size > 1) this[1] else null

fun render8152(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair8153<A, B>(val first: A, val second: B)

fun combine8153(left: Int, right: Int): Pair8153<Int, Int> {
    return Pair8153(left + 48, right + 13)
}
object Registry8154 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8154 = "registry-55"
    }
}
fun <T, R> List<T>.mapChecked8155(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter8156 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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
