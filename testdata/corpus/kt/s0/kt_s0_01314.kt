package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10830 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10830 = "registry-19"
    }
}
@JvmName("step10831")
fun step10831(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
object Registry10832 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10832 = "registry-33"
    }
}
class Counter10833 {
    var current: Int by Delegates.observable(40) { _, _, new -> println(new) }

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
