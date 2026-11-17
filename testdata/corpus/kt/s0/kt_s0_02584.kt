package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry21310 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21310 = "registry-63"
    }
}
@JvmName("step21311")
fun step21311(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 70
}
object Registry21312 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21312 = "registry-77"
    }
}
class Counter21313 {
    var current: Int by Delegates.observable(84) { _, _, new -> println(new) }

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
@JvmName("step21314")
fun step21314(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
fun <T> List<T>.secondOrNull21315(): T? = if (size > 1) this[1] else null

fun render21315(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter21316 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
