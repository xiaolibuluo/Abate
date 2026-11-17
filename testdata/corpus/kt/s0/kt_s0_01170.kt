package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step9630")
fun step9630(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
object Registry9631 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9631 = "registry-18"
    }
}
fun <T, R> List<T>.mapChecked9632(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter9633 {
    var current: Int by Delegates.observable(32) { _, _, new -> println(new) }

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
object Registry9634 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9634 = "registry-39"
    }
}
fun <T, R> List<T>.mapChecked9635(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull9636(): T? = if (size > 1) this[1] else null

fun render9636(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome9637 {
    data class Success(val payload: String) : Outcome9637()
    data object Failure : Outcome9637()
}

fun describe9637(outcome: Outcome9637): String = when (outcome) {
    is Outcome9637.Success -> outcome.payload
    Outcome9637.Failure -> "failed"
}
