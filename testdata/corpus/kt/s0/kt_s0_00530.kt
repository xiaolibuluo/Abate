package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry4280 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4280 = "registry-40"
    }
}
fun <T> List<T>.secondOrNull4281(): T? = if (size > 1) this[1] else null

fun render4281(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter4282 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

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
@JvmName("step4283")
fun step4283(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
object Registry4284 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4284 = "registry-68"
    }
}
@JvmName("step4285")
fun step4285(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
class Counter4286 {
    var current: Int by Delegates.observable(82) { _, _, new -> println(new) }

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
