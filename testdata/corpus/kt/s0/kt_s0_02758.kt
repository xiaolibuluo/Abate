package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter22770 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
@JvmName("step22771")
fun step22771(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
data class Pair22772<A, B>(val first: A, val second: B)

fun combine22772(left: Int, right: Int): Pair22772<Int, Int> {
    return Pair22772(left + 35, right + 49)
}
object Registry22773 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22773 = "registry-42"
    }
}
object Registry22774 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22774 = "registry-49"
    }
}
data class Pair22775<A, B>(val first: A, val second: B)

fun combine22775(left: Int, right: Int): Pair22775<Int, Int> {
    return Pair22775(left + 56, right + 21)
}
class Counter22776 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
