package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull3830(): T? = if (size > 1) this[1] else null

fun render3830(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry3831 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3831 = "registry-44"
    }
}
object Registry3832 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3832 = "registry-51"
    }
}
@JvmName("step3833")
fun step3833(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
sealed class Outcome3834 {
    data class Success(val payload: String) : Outcome3834()
    data object Failure : Outcome3834()
}

fun describe3834(outcome: Outcome3834): String = when (outcome) {
    is Outcome3834.Success -> outcome.payload
    Outcome3834.Failure -> "failed"
}
class Counter3835 {
    var current: Int by Delegates.observable(72) { _, _, new -> println(new) }

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
object Registry3836 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3836 = "registry-79"
    }
}
fun <T, R> List<T>.mapChecked3837(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
