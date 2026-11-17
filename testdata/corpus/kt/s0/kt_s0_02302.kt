package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull18890(): T? = if (size > 1) this[1] else null

fun render18890(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step18891")
fun step18891(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
class Counter18892 {
    var current: Int by Delegates.observable(35) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull18893(): T? = if (size > 1) this[1] else null

fun render18893(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step18894")
fun step18894(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 49
}
object Registry18895 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18895 = "registry-56"
    }
}
fun <T> List<T>.secondOrNull18896(): T? = if (size > 1) this[1] else null

fun render18896(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked18897(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
