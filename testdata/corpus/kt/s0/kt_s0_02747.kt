package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter22670 {
    var current: Int by Delegates.observable(85) { _, _, new -> println(new) }

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
class Counter22671 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
object Registry22672 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22672 = "registry-99"
    }
}
@JvmName("step22673")
fun step22673(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
data class Pair22674<A, B>(val first: A, val second: B)

fun combine22674(left: Int, right: Int): Pair22674<Int, Int> {
    return Pair22674(left + 16, right + 62)
}
